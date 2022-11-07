package go_render

import (
	// "log"

	"bytes"
	"go_render/pkg/gen"
	"image"
	"io"
	"log"

	"image/color"
	"image/png"
	_ "image/png"

	"golang.org/x/net/context"
)

type RenderInterface struct {
	gen.UnimplementedRenderInterfaceServer
	bgColourUpdate chan string
	sceneUpdate    chan string
}

func NewRenderInterface() *RenderInterface {

	bgColourUpdate := make(chan string)
	sceneUpdate := make(chan string)

	go func() {
		log.Printf("start bg channel")
		for {
			newBgColour := <-bgColourUpdate
			log.Printf("new colour received: %v", newBgColour)
		}
	}()

	go func() {
		log.Printf("start scene channel")
		for {
			newSceneUpdate := <-sceneUpdate
			log.Printf("new scene update: %v", newSceneUpdate)
		}
	}()

	return &RenderInterface{
		UnimplementedRenderInterfaceServer: gen.UnimplementedRenderInterfaceServer{},
		bgColourUpdate:                     bgColourUpdate,
		sceneUpdate:                        sceneUpdate,
	}
}

func (r *RenderInterface) SetBackgroundColour(ctx context.Context, in *gen.SetBackgroundColourRequest) (*gen.EmptyResponse, error) {

	r.bgColourUpdate <- in.GetColour()

	return &gen.EmptyResponse{}, nil
}

func (r *RenderInterface) Render(in *gen.RenderRequest, stream gen.RenderInterface_RenderServer) error {
	//bufferSize := 64 *1024 //64KiB, tweak this as desired

	newImage := image.NewNRGBA(image.Rect(0, 0, 100, 120))
	newImage.Set(10, 10, color.NRGBA{200,50,50,255})
	buf := bytes.NewBuffer([]byte{})

	err := png.Encode(buf, newImage)

	if err != nil {
		if err != io.EOF {
			log.Print(err)
		}
		return err
	}

	resp := &gen.RenderResponse{
		FileChunk: buf.Bytes(),
	}
	err = stream.Send(resp)
	if err != nil {
		log.Println("error while sending chunk:", err)
		return err
	}

	return nil
}
