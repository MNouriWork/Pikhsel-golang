package ui

func Setup(app *AppInit) {

	swatchContainer := BuildSwatches(app)

	app.PikhselWindow.SetContent(swatchContainer)

}
