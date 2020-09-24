# Organization scanner

Scan your repositories looking for hardcoded credentials (passwords, token). All leaks are collected and exported as a 
`.csv` file.

## Supported operations

- Scan all repositories of a given organization
    - `go run . -organization=<organization-name>`
    
- Scan single repository
    - `go run . -repository-url=<repository-ssh-url>`
    
## Scanning private repositories
To allow read your repository information, you must set a 
[GitHub token](https://docs.github.com/es/github/authenticating-to-github/creating-a-personal-access-token) in the 
[.env](.env) file. To scan the repos, the application clones the entire repository in memory, so you will need to 
configure an SSH key to allow cloning your repository. This key must be located in `$HOME/.ssh/id_rsa` (default). 
Learn more about how to configure SSH access 
[here](https://docs.github.com/es/github/authenticating-to-github/connecting-to-github-with-ssh).

### Scanner configuration
If you need to add custom regex expressions to detect secrets, just edit the [configuration file](scan-config.toml) 
adding/removing rules. You can find more examples about how rules are composed at 
[GitLeaks configuration docs](https://github.com/zricethezav/gitleaks/wiki/Configuration).

## Support
This project uses [go-github](https://github.com/google/go-github) to get repository information, to perform the scan 
uses [GitLeaks](https://github.com/zricethezav/gitleaks) as an API. It doesn't execute any command in your terminal, it 
performs the scanning using the library public methods.

## WishList
- ~~Scan single repository~~
- Centralize authentication strategy
- Flexible report output (support for `.json` as well)
- Execute inside Docker


