# Flickr Service

A simple JSON micro-service for providing photos from [flickr.com][flickr].

## Configuration

Configuraton is handled through the following environment varibles:

### Flickr
| Environment Variable |                           Description                           | Default Value |
|----------------------|-----------------------------------------------------------------|---------------|
| `FLICKR_ALBUM`       | The id of the flickr album to use as the source. (**Required**) |               |
| `FLICKR_KEY`         | A valid Flickr ID Key.(**Required**)                            |               |
| `REFRESH_INTERVAL`   | The length of time _in minutes_ between data updats.            | `15`          |

### API Configuration
|  Environment Variable |         Description         | Default Value |
|-----------------------|-----------------------------|---------------|
| `HOST`                | The host to listen on.      | `""` (all)    |
| `PORT`                | The HTTP port to listen on. | `3000`        |


## Run

```shell
env PORT=9000 \
    FLICKR_ALBUM=72157622339688060 \
    FLICKR_KEY=aa3e0e79a1a7bf4ae70290b9bdc98ee7 \
    flickr-service
```

[flickr]: https://www.flickr.com