// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package test

import (
	"github.com/goburrow/modbus"
	"log"
	"os"
	"testing"
	"time"
)

const (
	rtuDevice = "COM4"
)

//func TestRTUEncoding(t *testing.T) {
//	handler := modbus.NewRTUClientHandler(rtuDevice)
//	handler.BaudRate = 9600
//	handler.DataBits = 8
//	handler.Parity = "E"
//	handler.StopBits = 1
//	handler.SlaveId = 1
//	handler.Logger = log.New(os.Stdout, "rtu: ", log.LstdFlags)
//
//	pdu := modbus.ProtocolDataUnit{}
//	pdu.FunctionCode = 0x03
//	pdu.Data = []byte{0x50, 0x00, 0x00, 0x18}
//
//	adu, err2 := handler.Encode(&pdu)
//
//	err := handler.Connect()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer handler.Close()
//
//	client := modbus.NewClient(handler)
//	client.WriteRe
//	fmt.Println("client:", client)
//
//}

func TestRTUClient(t *testing.T) {
	// Diagslave does not support broadcast id.
	handler := modbus.NewRTUClientHandler(rtuDevice)
	handler.SlaveId = 17
	//ClientTestAll(t, modbus.NewClient(handler))
}

func TestRTUClientAdvancedUsage(t *testing.T) {
	handler := modbus.NewRTUClientHandler(rtuDevice)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Logger = log.New(os.Stdout, "rtu: ", log.LstdFlags)
	err := handler.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer handler.Close()

	client := modbus.NewClient(handler)

	//client.WriteSingleRegister(4, 0)
	//client.WriteSingleRegister(4, 0)

	// 加MODBUSCRC16校验，放到最后两位
	// 01 06 00 00 00 01
	results, err := client.WriteSingleRegister(17, 1)

	if err != nil || results == nil {
		t.Fatal(err, results)
	}
	time.Sleep(3 * time.Second)
	results, err = client.WriteSingleRegister(22, 1)
	//client.WriteSingleRegister(17, 1)

	//client.WriteSingleRegister(22, 1)

	//results, err = client.ReadDiscreteInputs(15, 2)
	//if err != nil || results == nil {
	//	t.Fatal(err, results)
	//}
	//results, err = client.ReadWriteMultipleRegisters(0, 2, 2, 2, []byte{1, 2, 3, 4})
	//if err != nil || results == nil {
	//	t.Fatal(err, results)
	//}
}
