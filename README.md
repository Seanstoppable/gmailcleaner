# Gmail Cleaner

A project to automate gmail inbox cleanup.

Inspired by [inboxbot](https://github.com/paulfurley/inboxbot).

Distinction is that I really just want to remove the inbox label to archive the
mail, and/or eventually apply other labels via rules.

Why not just use Gmail's filters? Gmail doesn't provide time based filters, it
only does filtering on new email ingestion.

## Setup

Note, you will have to follow the setup in the [go
quickstart](https://developers.google.com/gmail/api/quickstart/go) to set up
developer access and configure oauth.

Then, copy rules.yml.example to ~/.config/rules/rules.yml and edit as you see fit.

Finally, run ./gmailcleaner, follow the prompts, and you should start filtering
emails.
