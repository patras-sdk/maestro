package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/kudobuilder/kudo/pkg/kudoctl/cmd/env"
	"github.com/kudobuilder/kudo/pkg/kudoctl/cmd/env/kudohome"
	manInit "github.com/kudobuilder/kudo/pkg/kudoctl/cmd/init"
	"github.com/kudobuilder/kudo/pkg/kudoctl/kube"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

const initDesc = `
This command installs KUDO Manager onto your Kubernetes Cluster and sets up local configuration in $KUDO_HOME (default ~/.kudo/).

As with the rest of the KUDO commands, 'kudo init' discovers Kubernetes clusters
by reading $KUBECONFIG (default '~/.kube/config') and using the default context.

To set up just a local environment, use '--client-only'. That will configure
$KUDO_HOME, but not attempt to connect to a Kubernetes cluster and install the KUDO Manager.

When installing KUDO manager, 'kudo init' will attempt to install the latest released
version. You can specify an alternative image with '--kudo-image'.

To dump a manifest containing the KUDO deployment YAML, combine the
'--dry-run' and '--debug' flags.
`

type initCmd struct {
	out        io.Writer
	fs         afero.Fs
	image      string
	clientOnly bool
	dryRun     bool
	home       kudohome.Home
	wait       bool
}

func newInitCmd(fs afero.Fs, out io.Writer) *cobra.Command {
	i := &initCmd{fs: fs, out: out}

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize KUDO on both client and server",
		Long:  initDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				return errors.New("this command does not accept arguments")
			}
			i.home = Settings.Home

			return i.run()
		},
	}

	f := cmd.Flags()
	f.StringVarP(&i.image, "kudo-image", "i", "", "Override KUDO manager image and/or version")
	f.BoolVarP(&i.clientOnly, "client-only", "c", false, "If set does not install KUDO manager")
	f.BoolVar(&i.dryRun, "dry-run", false, "Do not install local or remote")
	f.BoolVar(&i.wait, "wait", false, "Block until KUDO manager is running and ready to receive requests")

	return cmd
}

// run initializes local config and installs KUDO manager to Kubernetes cluster.
func (i *initCmd) run() error {
	opts := manInit.NewOptions()
	// if image provide swtich to it.
	if i.image != "" {
		opts.Image = i.image
	}

	if Settings.Debug {
		mans, err := manInit.PrereqManifests(opts)
		if err != nil {
			return err
		}

		crd, err := manInit.CrdManifests()
		if err != nil {
			return err
		}

		deploy, err := manInit.DeploymentManifests(opts)
		if err != nil {
			return err
		}

		mans = append(mans, crd...)
		mans = append(mans, deploy...)
		i.YAMLWriter(i.out, mans)
	}
	if i.dryRun {
		return nil
	}

	// initialize client
	if err := initialize(i.home, i.out, Settings); err != nil {
		return fmt.Errorf("error initializing: %s", err)
	}
	fmt.Fprintf(i.out, "$KUDO_HOME has been configured at %s.\n", Settings.Home)

	// initialize server
	if !i.clientOnly {
		client, err := kube.GetKubeClient(Settings.KubeConfig)
		if err != nil {
			return fmt.Errorf("could not get kubernetes client: %s", err)
		}
		extClient, err := kube.GetKubeAPIExtClient(Settings.KubeConfig)
		if err != nil {
			return fmt.Errorf("could not get kubernetes extension client: %s", err)
		}

		if err := manInit.Install(client, extClient, opts); err != nil {
			if apierrors.IsAlreadyExists(err) {
				fmt.Fprintln(i.out, "Warning: KUDO manager is already installed in the cluster.\n"+
					"(Use --client-only to suppress this message)")
			} else {
				return fmt.Errorf("error installing: %s", err)
			}
		}

		if i.wait {
			manInit.WatchKUDOUntilReady(client, opts, 300)
		}

	} else {
		fmt.Fprintln(i.out, "Not installing KUDO manager due to 'client-only' flag having been set")
	}

	return nil
}

// YAMLWriter writes yaml to writer.   Looked into using https://godoc.org/gopkg.in/yaml.v2#NewEncoder which
// looks like a better way, however the omitted JSON elements are encoded which results in a very verbose output.
//TODO: Write a Encoder util which uses the "sigs.k8s.io/yaml" library for marshalling
func (i *initCmd) YAMLWriter(w io.Writer, manifests []string) error {
	for _, manifest := range manifests {
		if _, err := fmt.Fprintln(w, "---"); err != nil {
			return err
		}

		if _, err := fmt.Fprintln(w, manifest); err != nil {
			return err
		}
	}

	// YAML ending document boundary marker
	_, err := fmt.Fprintln(w, "...")
	return err
}

func initialize(home kudohome.Home, w io.Writer, settings env.Settings) error {
	//TODO (kensipe): implement the client init
	// on another pr
	return nil
}
