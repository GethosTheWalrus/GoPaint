clear

echo "Removing container GoPaint..."
sudo docker rm -f GoPaint &> /dev/null

echo "Building image gopaint..."
sudo docker build --no-cache -t gopaint . #&> /dev/null

echo "Running container GoPaint from image gopaint..."
sudo docker run -ti --name=GoPaint gopaint
