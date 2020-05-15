package main

import (
	"github.com/boyter/lc/parsers"
	"github.com/urfave/cli"
	"os"
)

//go:generate go run scripts/include.go
func main() {
	//defer profile.Start(profile.CPUProfile).Stop()
	//see apache2.0 license which uses special value for it so this does not work example ./lc/xgboost

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = parsers.ToolName
	app.Version = parsers.ToolVersion
	app.Usage = "Check directory for licenses and list what license(s) a file is under"
	app.UsageText = "lc [global options] [DIRECTORY|FILE] [DIRECTORY|FILE]"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "format, f",
			Usage:       "Set output format, supports progress, tabular, json, spdx, xlsx or `csv`",
			Destination: &parsers.Format,
			Value:       "tabular",
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "Set output file if not set will print to stdout `FILE`",
			Destination: &parsers.FileOutput,
		},
		cli.IntFlag{
			Name:        "filesize, fs",
			Usage:       "How large a file in bytes should be processed `50000`",
			Value:       50000,
			Destination: &parsers.MaxSize,
		},
		cli.StringFlag{
			Name:        "licensefiles, lf",
			Usage:       "Possible license files to inspect for over-arching license as comma separated list `copying,readme`",
			Value:       "license,licence,lisence,lisense,copying,readme,legal,copyright,copyleft,gpl,bsd,mit,apache,unlicense,unlicence",
			Destination: &parsers.PossibleLicenceFiles,
		},
		cli.StringFlag{
			Name:        "pathblacklist, pbl",
			Usage:       "Which directories should be ignored as comma seperated list `.git,.hg,.svn`",
			Value:       ".git,.hg,.svn",
			Destination: &parsers.PathBlacklist,
		},
		cli.StringFlag{
			Name:        "extblacklist, xbl",
			Usage:       "Which file extensions should be ignored for deep analysis as comma separated list E.G. `gif,jpg,png`",
			Value:       "woff,eot,cur,dm,xpm,emz,db,scc,idx,mpp,dot,pspimage,stl,dml,wmf,rvm,resources,tlb,docx,doc,xls,xlsx,ppt,pptx,msg,vsd,chm,fm,book,dgn,blines,cab,lib,obj,jar,pdb,dll,bin,out,elf,so,msi,nupkg,pyc,ttf,woff2,jpg,jpeg,png,gif,bmp,psd,tif,tiff,yuv,ico,xls,xlsx,pdb,pdf,apk,com,exe,bz2,7z,tgz,rar,gz,zip,zipx,tar,rpm,bin,dmg,iso,vcd,mp3,flac,wma,wav,mid,m4a,3gp,flv,mov,mp4,mpg,rm,wmv,avi,m4v,sqlite,class,rlib,ncb,suo,opt,o,os,pch,pbm,pnm,ppm,pyd,pyo,raw,uyv,uyvy,xlsm,swf",
			Destination: &parsers.ExtentionBlacklist,
		},
		cli.StringFlag{
			Name:        "documentname, dn",
			Usage:       "SPDX only. Sets DocumentName E.G. `LicenseChecker`",
			Value:       "Unknown",
			Destination: &parsers.DocumentName,
		},
		cli.StringFlag{
			Name:        "packagename, pn",
			Usage:       "SPDX only. Sets PackageName E.G. `LicenseChecker`",
			Value:       "Unknown",
			Destination: &parsers.PackageName,
		},
		cli.StringFlag{
			Name:        "documentnamespace, dns",
			Usage:       "SPDX only. Sets DocumentNamespace, if not set will default to http://spdx.org/spdxdocs/[packagename]-[HASH]",
			Destination: &parsers.DocumentNamespace,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Set to enable debug output",
			Destination: &parsers.Debug,
		},
		cli.BoolFlag{
			Name:        "trace",
			Usage:       "Set to enable debug output",
			Destination: &parsers.Trace,
		},
	}
	app.Action = func(c *cli.Context) error {
		parsers.DirFilePaths = c.Args()
		parsers.Process()
		return nil
	}

	app.Run(os.Args)
}
