package bunny

import "context"

// VideoLibraries represents the response of the List Video Library API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/videolibrarypublic_index
type VideoLibraries PaginationReply[VideoLibrary]

// List retrieves the Video Libraries.
// If opts is nil, DefaultPaginationPerPage and DefaultPaginationPage will be used.
// if opts.Page or or opts.PerPage is < 1, the related DefaultPagination values are used.
//
// Bunny.net API docs: https://docs.bunny.net/reference/videolibrarypublic_index
// TODO: add `includeAccessKey` path param
func (s *VideoLibraryService) List(
	ctx context.Context,
	opts *PaginationOptions,
) (*VideoLibraries, error) {
	return resourceList[VideoLibraries](ctx, s.client, "/videolibrary", opts)
}
