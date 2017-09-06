# delete service if it already exists
if (Get-Service jolokiaperfbeat -ErrorAction SilentlyContinue) {
  $service = Get-WmiObject -Class Win32_Service -Filter "name='jolokiaperfbeat'"
  $service.StopService()
  Start-Sleep -s 1
  $service.delete()
}

$workdir = Split-Path $MyInvocation.MyCommand.Path

# create new service
New-Service -name jolokiaperfbeat `
  -displayName jolokiaperfbeat `
  -binaryPathName "`"$workdir\\jolokiaperfbeat.exe`" -c `"$workdir\\jolokiaperfbeat.yml`" -path.home `"$workdir`"
