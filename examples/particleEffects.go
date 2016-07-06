package examples

import (
	"image/color"

	"github.com/walesey/go-engine/actor"
	"github.com/walesey/go-engine/assets"
	"github.com/walesey/go-engine/controller"
	"github.com/walesey/go-engine/effects"
	"github.com/walesey/go-engine/glfwController"
	"github.com/walesey/go-engine/opengl"
	"github.com/walesey/go-engine/renderer"
	"github.com/walesey/go-engine/vectormath"

	"github.com/codegangsta/cli"
)

//
func Particles(c *cli.Context) {
	fps := renderer.CreateFPSMeter(1.0)
	fps.FpsCap = 60

	glRenderer := &opengl.OpenglRenderer{
		WindowTitle:  "GoEngine",
		WindowWidth:  1400,
		WindowHeight: 900,
	}

	//setup scenegraph

	geom, _ := assets.ImportObj("TestAssets/Files/skybox/skybox.obj")
	geom.Material.LightingMode = renderer.MODE_UNLIT
	geom.CullBackface = false
	skyNode := renderer.CreateNode()
	skyNode.Add(geom)
	skyNode.SetRotation(1.57, vectormath.Vector3{0, 1, 0})
	skyNode.SetScale(vectormath.Vector3{5000, 5000, 5000})

	geomsphere, _ := assets.ImportObj("TestAssets/Files/sphere/sphere.obj")
	sphereNode := renderer.CreateNode()
	sphereNode.Add(geomsphere)
	sphereNode.SetTranslation(vectormath.Vector3{0, 0, 0})

	//particle effects
	explostionImg, _ := assets.ImportImage("TestAssets/Explosion.png")
	explosionMat := assets.CreateMaterial(explostionImg, nil, nil, nil)
	explosionMat.LightingMode = renderer.MODE_UNLIT
	explosionMat.Transparency = renderer.TRANSPARENCY_EMISSIVE
	explosionMat.DepthMask = false
	explosionParticles := effects.CreateParticleSystem(effects.ParticleSettings{
		MaxParticles:        4,
		ParticleEmitRate:    2,
		BaseGeometry:        renderer.CreateBox(float32(1), float32(1)),
		Material:            explosionMat,
		TotalFrames:         36,
		FramesX:             6,
		FramesY:             6,
		MaxLife:             1.0,
		MinLife:             2.0,
		StartSize:           vectormath.Vector3{0.4, 0.4, 0.4},
		EndSize:             vectormath.Vector3{2.4, 2.4, 2.4},
		StartColor:          color.NRGBA{254, 254, 254, 254},
		EndColor:            color.NRGBA{254, 254, 254, 254},
		MinTranslation:      vectormath.Vector3{-0.1, -0.1, -0.1},
		MaxTranslation:      vectormath.Vector3{0.1, 0.1, 0.1},
		MaxStartVelocity:    vectormath.Vector3{0.2, 1.8, 0.2},
		MinStartVelocity:    vectormath.Vector3{-0.2, 2.5, -0.2},
		Acceleration:        vectormath.Vector3{0.0, 0.0, 0.0},
		MaxAngularVelocity:  vectormath.IdentityQuaternion(),
		MinAngularVelocity:  vectormath.IdentityQuaternion(),
		MaxRotationVelocity: 0.0,
		MinRotationVelocity: 0.0,
	})

	fireImg, _ := assets.ImportImage("TestAssets/Fire.png")
	fireMat := assets.CreateMaterial(fireImg, nil, nil, nil)
	fireMat.LightingMode = renderer.MODE_UNLIT
	fireMat.Transparency = renderer.TRANSPARENCY_EMISSIVE
	fireMat.DepthMask = false
	fireParticles := effects.CreateParticleSystem(effects.ParticleSettings{
		MaxParticles:        10,
		ParticleEmitRate:    2,
		BaseGeometry:        renderer.CreateBox(float32(1), float32(1)),
		Material:            fireMat,
		TotalFrames:         36,
		FramesX:             6,
		FramesY:             6,
		MaxLife:             1.0,
		MinLife:             1.3,
		StartSize:           vectormath.Vector3{1.0, 1.0, 1.0},
		EndSize:             vectormath.Vector3{1.7, 1.7, 1.7},
		StartColor:          color.NRGBA{254, 54, 0, 200},
		EndColor:            color.NRGBA{254, 100, 20, 50},
		MinTranslation:      vectormath.Vector3{-0.1, 0.1, -0.1},
		MaxTranslation:      vectormath.Vector3{0.1, 0.3, 0.1},
		MaxStartVelocity:    vectormath.Vector3{0.02, 0.02, 0.02},
		MinStartVelocity:    vectormath.Vector3{-0.02, -0.02, -0.02},
		Acceleration:        vectormath.Vector3{0.0, 0.0, 0.0},
		MaxAngularVelocity:  vectormath.IdentityQuaternion(),
		MinAngularVelocity:  vectormath.IdentityQuaternion(),
		MaxRotationVelocity: 0.3,
		MinRotationVelocity: -0.3,
	})

	smokeImg, _ := assets.ImportImage("TestAssets/Smoke.png")
	smokeMat := assets.CreateMaterial(smokeImg, nil, nil, nil)
	smokeMat.LightingMode = renderer.MODE_UNLIT
	smokeMat.DepthMask = false
	smokeParticles := effects.CreateParticleSystem(effects.ParticleSettings{
		MaxParticles:        38,
		ParticleEmitRate:    15,
		BaseGeometry:        renderer.CreateBox(float32(1), float32(1)),
		Material:            smokeMat,
		TotalFrames:         64,
		FramesX:             8,
		FramesY:             8,
		MaxLife:             2.5,
		MinLife:             2.3,
		StartSize:           vectormath.Vector3{0.4, 0.4, 0.4},
		EndSize:             vectormath.Vector3{2.4, 2.4, 2.4},
		StartColor:          color.NRGBA{254, 254, 254, 50},
		EndColor:            color.NRGBA{254, 254, 254, 0},
		MinTranslation:      vectormath.Vector3{-0.2, -0.2, -0.2},
		MaxTranslation:      vectormath.Vector3{0.2, 0.2, 0.2},
		MaxStartVelocity:    vectormath.Vector3{0.2, 0.8, 0.2},
		MinStartVelocity:    vectormath.Vector3{-0.2, 0.6, -0.2},
		Acceleration:        vectormath.Vector3{0.0, 0.0, 0.0},
		MaxAngularVelocity:  vectormath.IdentityQuaternion(),
		MinAngularVelocity:  vectormath.IdentityQuaternion(),
		MaxRotationVelocity: 0.0,
		MinRotationVelocity: 0.0,
	})

	sparkImg, _ := assets.ImportImage("TestAssets/test.png")
	sparkMat := assets.CreateMaterial(sparkImg, nil, nil, nil)
	sparkMat.LightingMode = renderer.MODE_EMIT
	sparkMat.Transparency = renderer.TRANSPARENCY_EMISSIVE
	sparkParticles := effects.CreateParticleSystem(effects.ParticleSettings{
		MaxParticles:        100,
		ParticleEmitRate:    110,
		BaseGeometry:        renderer.CreateBox(float32(1), float32(1)),
		Material:            sparkMat,
		TotalFrames:         1,
		FramesX:             1,
		FramesY:             1,
		MaxLife:             0.9,
		MinLife:             0.7,
		StartSize:           vectormath.Vector3{0.02, 0.02, 0.02},
		EndSize:             vectormath.Vector3{0.02, 0.02, 0.02},
		StartColor:          color.NRGBA{255, 5, 5, 255},
		EndColor:            color.NRGBA{255, 5, 5, 255},
		MinTranslation:      vectormath.Vector3{0, -0, 0},
		MaxTranslation:      vectormath.Vector3{0, -0, 0},
		MaxStartVelocity:    vectormath.Vector3{0.6, 0.3, 0.6},
		MinStartVelocity:    vectormath.Vector3{-0.6, 0.5, -0.6},
		Acceleration:        vectormath.Vector3{0.0, 0.0, 0.0},
		MaxAngularVelocity:  vectormath.IdentityQuaternion(),
		MinAngularVelocity:  vectormath.IdentityQuaternion(),
		MaxRotationVelocity: 0.0,
		MinRotationVelocity: 0.0,
	})
	fireParticles.Location = vectormath.Vector3{2, 0, -2}
	smokeParticles.Location = vectormath.Vector3{-2, 0, 2}
	explosionParticles.Location = vectormath.Vector3{-2, 0, -2}
	sparkParticles.Location = vectormath.Vector3{2, 0, 2}

	birdImg, _ := assets.ImportImage("TestAssets/test.png")
	birdMat := assets.CreateMaterial(birdImg, nil, nil, nil)
	birdMat.LightingMode = renderer.MODE_UNLIT
	birdSprite := effects.CreateSprite(22, 5, 5, birdMat)
	birdSprite.SetTranslation(vectormath.Vector3{-2, 0, -1})

	sceneGraph := renderer.CreateSceneGraph()
	sceneGraph.Add(skyNode)
	sceneGraph.Add(sphereNode)
	sceneGraph.AddTransparent(fireParticles)
	sceneGraph.AddTransparent(smokeParticles)
	sceneGraph.AddTransparent(explosionParticles)
	sceneGraph.AddTransparent(birdSprite)
	sceneGraph.AddTransparent(sparkParticles)

	//camera
	camera := renderer.CreateCamera(glRenderer)
	freeMoveActor := actor.NewFreeMoveActor(camera)
	freeMoveActor.MoveSpeed = 3.0
	freeMoveActor.Location = vectormath.Vector3{-2, 0, 0}

	glRenderer.Init(func() {
		//Lighting
		glRenderer.CreateLight(0.1, 0.1, 0.1, 1, 1, 1, 1, 1, 1, true, vectormath.Vector3{0, -1, 0}, 0)

		//setup reflection map
		cubeMap := renderer.CreateCubemap(geom.Material.Diffuse)
		glRenderer.ReflectionMap(cubeMap)

		//post effects
		// cell := renderer.Shader{
		// 	Name: "shaders/cell/cellCoarse",
		// }
		// bloomHorizontal := renderer.Shader{
		// 	Name: "shaders/bloom/bloomHorizontal",
		// 	Uniforms: []renderer.Uniform{
		// 		renderer.Uniform{"size", mgl32.Vec2{1900, 1000}},
		// 		renderer.Uniform{"quality", 2.5},
		// 		renderer.Uniform{"samples", 12},
		// 		renderer.Uniform{"threshold", 0.995},
		// 		renderer.Uniform{"intensity", 1.9},
		// 	},
		// }
		// bloomVertical := renderer.Shader{
		// 	Name: "shaders/bloom/bloomVertical",
		// 	Uniforms: []renderer.Uniform{
		// 		renderer.Uniform{"size", mgl32.Vec2{1900, 1000}},
		// 		renderer.Uniform{"quality", 2.5},
		// 		renderer.Uniform{"samples", 12},
		// 		renderer.Uniform{"threshold", 0.995},
		// 		renderer.Uniform{"intensity", 1.9},
		// 	},
		// }

		//input/controller manager
		controllerManager := glfwController.NewControllerManager(glRenderer.Window)

		//lock the cursor
		glRenderer.LockCursor(true)

		//camera free move actor
		mainController := controller.NewBasicMovementController(freeMoveActor, false)
		controllerManager.AddController(mainController.(glfwController.Controller))

		customController := controller.CreateController()
		controllerManager.AddController(customController.(glfwController.Controller))

		//close window and exit on escape
		customController.BindAction(func() {
			glRenderer.Window.SetShouldClose(true)
		}, controller.KeyEscape, controller.Press)

		//test the portabitity of the actor / entity interfaces
		customController.BindAction(func() { freeMoveActor.Entity = camera }, controller.KeyQ, controller.Press)
		customController.BindAction(func() { freeMoveActor.Entity = sphereNode }, controller.KeyW, controller.Press)
		customController.BindAction(func() { freeMoveActor.Entity = explosionParticles }, controller.KeyE, controller.Press)
		customController.BindAction(func() { freeMoveActor.Entity = birdSprite }, controller.KeyR, controller.Press)

		// customController.BindAction(func() { //no post effects
		// 	glRenderer.DestroyPostEffects(bloomVertical)
		// 	glRenderer.DestroyPostEffects(bloomHorizontal)
		// 	glRenderer.DestroyPostEffects(cell)
		// }, controller.KeyA, controller.Press)
		//
		// customController.BindAction(func() { //bloom effect
		// 	glRenderer.CreatePostEffect(bloomVertical)
		// 	glRenderer.CreatePostEffect(bloomHorizontal)
		// 	glRenderer.DestroyPostEffects(cell)
		// }, controller.KeyS, controller.Press)
		//
		// customController.BindAction(func() { //cell effect
		// 	glRenderer.DestroyPostEffects(bloomVertical)
		// 	glRenderer.DestroyPostEffects(bloomHorizontal)
		// 	glRenderer.CreatePostEffect(cell)
		// }, controller.KeyD, controller.Press)
	})

	glRenderer.Update(func() {
		fps.UpdateFPSMeter()

		//update things that need updating
		explosionParticles.SetCameraLocation(glRenderer.CameraLocation())
		fireParticles.SetCameraLocation(glRenderer.CameraLocation())
		smokeParticles.SetCameraLocation(glRenderer.CameraLocation())
		sparkParticles.SetCameraLocation(glRenderer.CameraLocation())
		explosionParticles.Update(0.018)
		fireParticles.Update(0.018)
		smokeParticles.Update(0.018)
		sparkParticles.Update(0.018)

		birdSprite.NextFrame()

		freeMoveActor.Update(0.018)
	})

	glRenderer.Render(func() {
		sceneGraph.RenderScene(glRenderer)
	})

	glRenderer.Start()
}
