# Omnifocus CLI

> Quickly send items to your Omnifocus inbox

<p>
    <a href="https://godoc.org/github.com/eugenetriguba/of">
        <img src="https://godoc.org/github.com/eugenetriguba/of?status.svg" alt="GoDoc">
    </a>
    <a href="https://goreportcard.com/report/github.com/eugenetriguba/of">
        <img src="https://goreportcard.com/badge/github.com/eugenetriguba/of" alt="Go Report Card Badge">
    </a>
    <a href="https://codebeat.co/projects/github-com-eugenetriguba-of-master">
        <img alt="codebeat badge" src="https://codebeat.co/badges/4bfb8156-c136-4544-bbe7-f5a842e4594c" />
    </a>
    <img alt="Version Badge" src="https://img.shields.io/badge/version-0.2.0-blue" style="max-width:100%;">
</p>

The Omnifocus CLI allows you to quickly send a todo to your Omnifocus
inbox. It makes use of [Omnifocus's mail drop](https://support.omnigroup.com/omnifocus-mail-drop/) feature. This feature gives you an email address that you can send emails to. When an email is sent to this address, it uses that email to create a todo in your Omnifocus inbox. This CLI
allows you to quickly send that email using an ``of`` command.

Before you can start sending todos into your inbox, you'll have to setup a few pieces of 
configuration. 
```bash
// Sets your Omnifocus mail drop email
of config --maildrop fake-mailbox@omnisync.com

// Sets your gmail username
of config --username email@gmail.com

// Sets your gmail password
of config --password secretpassword
```

Once configured, you can add todos into your inbox.
```bash
$ of add "Go to the grocery store" --note "Make sure to get milk" --attachment "~/report.pdf"
  Successfully sent your todo!
```
