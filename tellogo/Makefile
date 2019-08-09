PACKAGE=github.com/numa08/tello/tellogo

.PHONY: all

all: android ios

android:
	mkdir -p ../app/android/tellogo/
	gomobile bind -o ../app/android/tellogo/tellogo.aar --target=android $(PACKAGE)

ios:
	gomobile bind -o ../app/ios/Flutter/Tellogo.framework --target=ios $(PACKAGE)

clean:
	@rm -rf articles
	gomobile clean
