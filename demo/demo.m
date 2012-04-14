// clang demo.m -framework Foundation -framework AppKit -o demo-clang

#import <Foundation/Foundation.h>
#import <AppKit/AppKit.h>

int main() {
	NSAutoreleasePool *pool = [[NSAutoreleasePool alloc] init];

	NSApplication *app = [NSApplication sharedApplication];

	NSWindow *win = [[NSWindow alloc] initWithContentRect:CGRectMake(0, 0, 500, 500)
						styleMask:0 backing:NSBackingStoreBuffered defer:false];
	[win autorelease];
	[win display];	
	[win makeKeyAndOrderFront:win];
	[win setTitle:@"Go Demo"];

	[app run];

	[pool release];

	return 0;
}
