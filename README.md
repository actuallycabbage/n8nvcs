# n8n-vcs

Easily migrate your [n8n](https://github.com/n8n-io/n8n) configs between environments. `n8n-vcs pull` and `n8n-vcs push`

## Background
When working with multiple environments (or version control in general) you need some features like
- automatically read configs
- automatically deploy configs
- config change history
- upgrade/downgrade to any config in the history

n8n doesn't really have any of those features. n8n-vcs intends to (externally) add in functionality like that
