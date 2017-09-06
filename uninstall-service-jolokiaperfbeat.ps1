# delete service if it exists
if (Get-Service jolokiaperfbeat -ErrorAction SilentlyContinue) {
  $service = Get-WmiObject -Class Win32_Service -Filter "name='jolokiaperfbeat'"
  $service.delete()
}
