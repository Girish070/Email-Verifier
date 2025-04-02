# Domain Email Security Checker

This Go program checks domain email security settings by querying DNS records for MX, SPF, and DMARC information.

## Features
- Checks if a domain has MX (Mail Exchange) records.
- Looks for SPF (Sender Policy Framework) records.
- Checks for DMARC (Domain-based Message Authentication, Reporting, and Conformance) records.

## Prerequisites
- Go installed on your system ([Download Go](https://go.dev/dl/)).
- Internet connection to query DNS records.

## Installation
Clone the repository and navigate to the project directory:
```sh
git clone <your-repo-url>
cd <your-project-directory>
```

## Usage
Compile and run the program:
```sh
go run main.go
```

Then, enter domain names one per line in the terminal. Press `Ctrl+D` (Linux/macOS) or `Ctrl+Z` (Windows) to stop input.

### Example Output
```
domain, hasMX, hasSpf, spfRecord, hasDmarc, dmarcRecord
example.com, true, true, v=spf1 include:_spf.google.com ~all, true, v=DMARC1; p=none; rua=mailto:dmarc@example.com
```

## Error Handling
- If a DNS lookup fails, an error message is printed, but execution continues.

## License
This project is licensed under the MIT License.

