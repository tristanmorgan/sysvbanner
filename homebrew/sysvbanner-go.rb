require 'formula'

SYSVBANNER_GO_VERSION = '0.1.2'

class SysvbannerGo < Formula
  homepage 'https://github.com/winebarrel/sysvbanner-go'
  url "https://github.com/winebarrel/sysvbanner-go/releases/download/#{SYSVBANNER_GO_VERSION}/sysvbanner-go-#{SYSVBANNER_GO_VERSION}-darwin-amd64.tar.gz"
  sha256 '1192c8cb86b11dc46c7c60ff70c01cf75f7648af5ce431753f2ecb2bc3aef655'
  version SYSVBANNER_GO_VERSION
  head 'https://github.com/winebarrel/sysvbanner-go.git', :branch => 'master'

  def install
    bin.install 'banner'
  end
end
