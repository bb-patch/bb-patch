package patch

var Patches = []PatchFactory{
	&noInstallPatchFactory{English},
	&noCdPatchFactory{},
	&directDrawFixFactory{},
	&developerConsolePatchFactory{},
	&setResolutionPatchFactory{1280, 720},
}
