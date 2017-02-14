require 'formula'

SYSVBANNER_GO_VERSION = '0.1.2'

class SysvbannerGo < Formula
  homepage 'https://github.com/winebarrel/sysvbanner-go'
  url "https://github.com/winebarrel/sysvbanner-go/releases/download/#{SYSVBANNER_GO_VERSION}/sysvbanner-go-#{SYSVBANNER_GO_VERSION}-darwin-amd64.tar.gz"
  sha256 '65e51fcbe4319b13248c14a1445176ccb194bb7323ddb8133f9266ce9255fa14'
  version SYSVBANNER_GO_VERSION
  head 'https://github.com/winebarrel/sysvbanner-go.git', :branch => 'master'

  def install
    bin.install 'banner'
  end
end
