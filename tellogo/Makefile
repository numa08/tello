PACKAGE=github.com/numa08/tello/tellogo

.PHONY: all

all: android ios

android:
	mkdir -p articles/android
	gomobile bind -o articles/android/tellogo.aar --target=android $(PACKAGE)

ios:
	mkdir -p articles/ios
	gomobile bind -o articles/ios/Tellogo.framework --target=ios $(PACKAGE)

clean:
	@rm -rf articles
	gomobile clean
