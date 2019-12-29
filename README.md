# Omnifocus CLI

> Quickly send items to your inbox

<p>
    <a href="https://goreportcard.com/report/github.com/eugenetriguba/of">
        <img src="https://goreportcard.com/badge/github.com/eugenetriguba/of">
    </a>
    <a href="https://codebeat.co/projects/github-com-eugenetriguba-of-master">
        <img alt="codebeat badge" src="https://codebeat.co/badges/4bfb8156-c136-4544-bbe7-f5a842e4594c" />
    </a>
</p>

Omnifocus is a todo application that is only on macOS and iOS. To combat that, you can use this cli tool to quickly send an item to your inbox when you're using a different device.

The interface for this application is quite simple. There are a few pieces of configuration: your omnifocus mail drop address, gmail email, and gmail password. 
From there, you can easily send items to your inbox with the task name along with an optional note or attachment.

Usage:
```bash
$ of -email fake.mail.drop.email@sync.omnigroup.com
  Successfully set your omnifocus mail drop email to fake.mail.drop.email@sync.omnigroup.com
$ of -note "Make sure he is on his leash" -attachment ~/dogs/walk-schedule.pdf "Take dog for a walk"
  Successfully sent your todo!
```
