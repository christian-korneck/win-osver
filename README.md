# osver (Windows)

Prints the Windows OS version. Intended for Windows container troubleshooting. (Uses the same code path as Docker/Moby to get the host's os version for matching with `os.version` in container image manifests).

## Usage

```text
osver.exe [-b]

  -b    print only the build version (i.e. '2200' instead of '10.0.22000')
```

## Examples

```bash
$ osver.exe
10.0.22000
$ osver.exe -b
22000
```
