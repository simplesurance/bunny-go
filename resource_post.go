package bunny

import "context"

func resourcePost[Resp any](
	ctx context.Context,
	client *Client,
	path string,
	opts any,
) (*Resp, error) {
	var res Resp

	req, err := client.newPostRequest(path, opts)
	if err != nil {
		return nil, err
	}

	if err := client.sendRequest(ctx, req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func resourcePostWith204Response(
	ctx context.Context,
	client *Client,
	path string,
	opts any,
) error {
	_, err := resourcePost[NoContentResponse](ctx, client, path, opts)
	return err
}
