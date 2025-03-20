# gato-cli

**gato-cli** is a simple cli application for image processing built with [**gato**](https://github.com/obzva/gato).

## Usage

### install

```bash
git clone https://github.com/obzva/gato-cli.git

cd ./gato-cli

go build
```

### Command Line Interface

```bash
# example
./gato-cli -w 800 -h 600 -m bilinear -o output.jpg input.jpg
```

### flags

- `-w`: Desired width of output image, defaults to keep the ratio of the original image when omitted (**at least one of two, width or height, is required**)
- `-h`: Desired height of output image, defaults to keep the ratio of the original image when omitted (**at least one of two, width or height, is required**)
- `-m`: Interpolation method, defaults to nearestneighbor when omitted (options: nearestneighbor, bilinear, bicubic)
- `-o`: Output filename, defaults to the method name when omitted
- `-c`: Concurrency mode, defaults to true when omitted

### Example I/O

Let's scale sample image up twice

![input image](https://raw.githubusercontent.com/obzva/assets/refs/heads/main/gato-cli/test-image.jpg)

input image (500 x 300)

![nearest neighbor output image](https://raw.githubusercontent.com/obzva/assets/refs/heads/main/gato-cli/nearestneighbor.jpg)

output image using nearest neighbor interpolation (1000 x 600)

![bilinear output image](https://raw.githubusercontent.com/obzva/assets/refs/heads/main/gato-cli/bilinear.jpg)

output image using bilinear interpolation (1000 x 600)

![bicubic output image](https://raw.githubusercontent.com/obzva/assets/refs/heads/main/gato-cli/bicubic.jpg)

output image using bicubic interpolation (1000 x 600)
