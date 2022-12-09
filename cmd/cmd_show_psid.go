package cmd

import (
	"GoSungrow/iSolarCloud"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
	"strings"
)


func (c *CmdShow) AttachPs(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		var self = &cobra.Command{
			Use:                   "ps",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "PsId"},
			Short:                 fmt.Sprintf("Ps related Sungrow commands."),
			Long:                  fmt.Sprintf("Ps related Sungrow commands."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(self)
		self.Example = cmdHelp.PrintExamples(self, "")

		c.AttachPsList(self)
		c.AttachPsTree(self)
		c.AttachPsPoints(self)
		c.AttachPsData(self)
		c.AttachPsGraph(self)
	}
	return c.SelfCmd
}


func (c *CmdShow) AttachPsList(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "list",
		Aliases:               []string{"ls"},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Show all devices on account."),
		Long:                  fmt.Sprintf("Show all devices on account."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcPsList,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcPsList(_ *cobra.Command, args []string) error {
	for range Only.Once {
		var devices string
		devices, c.Error = cmds.Api.SunGrow.PsList(args...)
		if c.Error != nil {
			break
		}

		fmt.Printf("%s\n", devices)
	}
	return c.Error
}

func (c *CmdShow) AttachPsIdList2(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "list2",
		Aliases:               []string{"ls"},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Show all available PS."),
		Long:                  fmt.Sprintf("Show all available PS."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachPsIdList2,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachPsIdList2(_ *cobra.Command, args []string) error {
	for range Only.Once {
		pids := cmds.Api.SunGrow.SetPsIds(args...)
		if c.Error != nil {
			break
		}

		fmt.Printf("%s\n", pids)
	}
	return c.Error
}

func (c *CmdShow) AttachPsTree(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "tree",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Show the PS tree."),
		Long:                  fmt.Sprintf("Show the PS tree."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachPsTree,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachPsTree(_ *cobra.Command, args []string) error {
	for range Only.Once {
		var pids iSolarCloud.PsTree
		pids, c.Error = cmds.Api.SunGrow.PsTreeMenu(args...)
		if c.Error != nil {
			break
		}

		fmt.Printf("%s\n", pids)
	}
	return c.Error
}

func (c *CmdShow) AttachPsPoints(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "points [ps_ids | .] [device_type]",
		Aliases:               []string{"point"},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("List points used for a given ps_id."),
		Long:                  fmt.Sprintf("List points used for a given ps_id."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachPsPoints,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachPsPoints(_ *cobra.Command, args []string) error {
	for range Only.Once {
		args = MinimumArraySize(2, args)
		var points string
		points, c.Error = cmds.Api.SunGrow.PsPoints(strings.Split(args[0], ","), args[1])
		if c.Error != nil {
			break
		}

		fmt.Printf("%s\n", points)
	}
	return c.Error
}

func (c *CmdShow) AttachPsData(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "data <ps_ids | .> [device_type | .] [start date] [end date] [interval]",
		Aliases:               []string{"point"},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Generate points table for a given ps_id."),
		Long:                  fmt.Sprintf("Generate points table for a given ps_id."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachPsData,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachPsData(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetTable()
		args = MinimumArraySize(5, args)
		c.Error = cmds.Api.SunGrow.PsPointsData(strings.Split(args[0], ","), args[1], args[2], args[3], args[4])
		if c.Error != nil {
			break
		}
	}
	return c.Error
}

func (c *CmdShow) AttachPsGraph(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "graph <ps_ids | .> [device_type]",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Generate graphs of points for a given ps_id."),
		Long:                  fmt.Sprintf("Generate graphs of points for a given ps_id."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachPsGraph,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachPsGraph(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetGraph()
		args = MinimumArraySize(5, args)
		c.Error = cmds.Api.SunGrow.PsPointsData(strings.Split(args[0], ","), args[1], args[2], args[3], args[4])
		if c.Error != nil {
			break
		}
	}
	return c.Error
}
