// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of KubeArmor

package cmd

import (
	"github.com/kubearmor/kubearmor-client/vm"
	"github.com/spf13/cobra"
)

var vmOptions vm.VmOptions

// vmCmd represents the vm command
var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Download VM install script from kvmsoperator service",
	Long:  `Download VM install script from kvmsoperator service`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := vm.FileDownload(client, vmOptions); err != nil {
			return err
		}
		return nil
	},
}

// ========== //
// == Init == //
// ========== //

func init() {
	rootCmd.AddCommand(vmCmd)
	vmCmd.Flags().StringVarP(&vmOptions.Port, "port", "p", "32770", "Port of kvmsoperator")
	vmCmd.Flags().StringVarP(&vmOptions.Namespace, "ns", "n", "kube-system", "Namespace of kvmsoperator")
	vmCmd.Flags().StringVarP(&vmOptions.VMName, "kvm", "v", "", "Name of configured vm")
	vmCmd.Flags().StringVarP(&vmOptions.File, "file", "f", "none", "Filename with path to store the configured vm installation script")

	// Marking this flag as markedFlag and mandatory
	vmCmd.MarkFlagRequired("kvm")
}