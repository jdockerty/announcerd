# announcerd

Announce a message from pull requests into Slack

### Motivation

I have forgotten to announce too many times at work. This workflow involved moving from a pull
request and writing out the message myself into a channel. This made me think *"shouldn't think be done entirely from a PR?"*.

By simply adding `announcement="<message>"` to a pull request, the given message should be placed into a channel for the provided webhook.

### Usage

This is built as a [GitHub App](https://docs.github.com/en/developers/apps/getting-started-with-apps/about-apps) and will require installation as such.

#### Configuration

The application takes various environment variables which are used for configuration.

* `ANNOUNCERD_GH_APP_ID` is the GitHub App Application ID (required).
* `ANNOUNCERD_GH_APP_KEY_FILE` is the path to the `.pem` key file for the GitHub App (required).
* `ANNOUNCERD_SLACK_WEBHOOK` is the "Incoming Webhook" for Slack, configured to the channel you want to send announcements too (required).
* `ANNOUNCERD_HOST` is the address which `announcerd` binds to, by default this is `localhost`.
* `ANNOUNCERD_PORT` is the port which `announcerd` binds to, by default this is `6000`.

