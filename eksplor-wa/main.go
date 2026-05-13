package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var client *whatsmeow.Client

func main() {

	// Database session
	dbLog := waLog.Stdout("Database", "INFO", true)

	container, err := sqlstore.New(
		context.Background(),
		"sqlite3",
		"file:session.db?_foreign_keys=on",
		dbLog,
	)

	if err != nil {
		panic(err)
	}

	// Ambil device
	deviceStore, err := container.GetFirstDevice(context.Background())

	if err != nil {
		panic(err)
	}

	// Buat client WhatsApp
	client = whatsmeow.NewClient(
		deviceStore,
		waLog.Stdout("Client", "INFO", true),
	)

	// Jika belum login
	if client.Store.ID == nil {

		qrChan, _ := client.GetQRChannel(context.Background())

		err = client.Connect()

		if err != nil {
			panic(err)
		}

		// Tampilkan QR
		for evt := range qrChan {

			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}

	} else {

		err = client.Connect()

		if err != nil {
			panic(err)
		}
	}

	// Nomor tujuan
	nomor := "6285624610171"

	// Membuat JID WhatsApp
	jid := types.NewJID(nomor, "s.whatsapp.net")

	// Isi pesan
	text := "Halo, pesan dari Golang"

	// Membuat message
	msg := &waProto.Message{
		Conversation: &text,
	}

	// Kirim pesan
	_, err = client.SendMessage(
		context.Background(),
		jid,
		msg,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Pesan berhasil dikirim")

	select {}
}
