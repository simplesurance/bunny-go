package bunny

import "context"

func resourceGet[Resp any](
	ctx context.Context,
	client *Client,
	path string,
	opts interface{},
) (*Resp, error) {
	var res Resp

	req, err := client.newGetRequest(path, opts)
	if err != nil {
		return nil, err
	}

	if err := client.sendRequest(ctx, req, &res); err != nil {
		return nil, err
	}

	return &res, err
}
