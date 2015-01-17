require 'formula'

SYSVBANNER_GO_VERSION = '0.1.2'

class SysvbannerGo < Formula
  homepage 'https://github.com/winebarrel/sysvbanner-go'
  url "https://github.com/winebarrel/sysvbanner-go/releases/download/#{SYSVBANNER_GO_VERSION}/sysvbanner-go-#{SYSVBANNER_GO_VERSION}-darwin-amd64.tar.gz"
  sha1 '2ce97081f053d603813d33be509af0813e0a2f97'
  version SYSVBANNER_GO_VERSION
  head 'https://github.com/winebarrel/sysvbanner-go.git', :branch => 'master'

  def install
    bin.install 'banner'
  end
end
