package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xiahuaxiahua0616/minehub/cmd/mh-apiserver/app/options"
	"github.com/xiahuaxiahua0616/minehub/internal/pkg/log"
	"github.com/xiahuaxiahua0616/minehub/pkg/version"
)

var configFile string

func NewMineHubCommand() *cobra.Command {
	opts := options.NewServerOptions()
	cmd := &cobra.Command{
		Use:          "how use minehub",
		Short:        "short minehub",
		Long:         "long minehub",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
		Args: cobra.NoArgs,
	}

	// 推荐使用配置文件来配置应用，便于管理配置项
	// cmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the miniblog configuration file.")
	version.AddFlags(cmd.PersistentFlags())
	// 将 ServerOptions 中的选项绑定到命令标志
	opts.AddFlags(cmd.PersistentFlags())

	cobra.OnInitialize(onInitialize)

	return cmd
}

func run(opts *options.ServerOptions) error {
	version.PrintAndExitIfRequested()

	// 初始化日志
	log.Init(logOptions())
	defer log.Sync()

	if err := viper.Unmarshal(opts); err != nil {
		return err
	}

	if err := opts.Validate(); err != nil {
		return err
	}

	cfg, err := opts.Config()
	if err != nil {
		return err
	}

	server, err := cfg.NewUnionServer()
	if err != nil {
		return err
	}

	return server.Run()
}

func logOptions() *log.Options {
	opts := log.NewOptions()
	if viper.IsSet("log.disable-caller") {
		opts.DisableCaller = viper.GetBool("log.disable-caller")
	}
	if viper.IsSet("log.disable-stacktrace") {
		opts.DisableStacktrace = viper.GetBool("log.disable-stacktrace")
	}
	if viper.IsSet("log.level") {
		opts.Level = viper.GetString("log.level")
	}
	if viper.IsSet("log.format") {
		opts.Format = viper.GetString("log.format")
	}
	if viper.IsSet("log.output-paths") {
		opts.OutputPaths = viper.GetStringSlice("log.output-paths")
	}
	return opts
}
