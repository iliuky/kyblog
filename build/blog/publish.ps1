$CurrentyDir = Split-Path -Parent $MyInvocation.MyCommand.Definition;
$MainDir = $ExecutionContext.SessionState.Path.GetUnresolvedProviderPathFromPSPath("../../cmd/blog")
$RootDir = $ExecutionContext.SessionState.Path.GetUnresolvedProviderPathFromPSPath("../../")
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$GoBin = $env:GOPATH
$GoBin = $GoBin.Trim(";")
go install $MainDir

#New-Item -Path $CurrentyDir -Name publish -Type Directory -force
copy-item $MainDir\appsettings.production.toml -destination $CurrentyDir\publish
copy-item $MainDir\appsettings.toml -destination $CurrentyDir\publish
copy-item $GoBin\bin\linux_amd64\* -destination $CurrentyDir\publish
copy-item $CurrentyDir\dockerfile -destination $CurrentyDir\publish
copy-item $RootDir\public -destination $CurrentyDir\publish -Recurse -force
copy-item $RootDir\templates -destination $CurrentyDir\publish -Recurse -force