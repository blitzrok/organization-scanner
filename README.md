# Organization scanner

Scan your repositories looking for hardcoded credentials (passwords, token). All leaks are collected and exported as a 
`.csv` file.

## Supported operations

- Scan all repositories of a given organization
    - `go run . -organization=<organization-name>`
    
- Scan single repository
    - `go run . -repository-url=<repository-ssh-url>`
    
## Scanning private repositories
To allow read your repository information, you must set a GitHub token in the [.env](.env) file. To scan the repos, 
the application clones the entire repository in memory, so you will need to configure your SSH key to allow cloning your 
repository. This key must be located in `$HOME/.ssh/id_rsa` (default). 

## Support
This project uses [go-github](https://github.com/google/go-github) to get repository information, to perform the scan 
uses [GitLeaks](https://github.com/zricethezav/gitleaks) as an API. It doesn't execute any command in your terminal, it 
performs the scanning using the library public methods.

## Wishlist
- ~~Scan single repository~~
- Centralize authentication strategy
- Flexible report output (support for `.json` as well)
- Execute inside Docker
