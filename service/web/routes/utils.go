package routes

import (
	// Any route packages that need to be initialized.
	_ "github.com/SumeruCCTV/sumeru/service/web/routes/r_auth"
	_ "github.com/SumeruCCTV/sumeru/service/web/routes/r_camera"
)

func Init() {}
