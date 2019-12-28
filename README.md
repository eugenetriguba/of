# Omnifocus CLI

    Quickly send items to your inbox

Omnifocus is a todo application that is only on macOS and iOS. To combat that, you can use this cli tool to quickly send an item to your inbox when you're using a different device.

The interface for this application is quite simple. There is one bit of configuration, your omnifocus mail drop email. From there, you can easily sent items to your inbox with the task name and an optional note or attachment.

Usage:
```bash
$ of -email fake.mail.drop.email@sync.omnigroup.com
  Successfully set your omnifocus mail drop email to fake.mail.drop.email@sync.omnigroup.com
$ of "Take dog for a walk" -note "Make sure he is on his leash" -attachment ~/dogs/walk-schedule.pdf
  Successfully sent your todo!
```