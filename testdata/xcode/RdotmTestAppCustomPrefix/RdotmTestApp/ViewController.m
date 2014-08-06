//
//  ViewController.m
//  RdotmTestApp
//
//  Copyright (c) 2014 Soichiro Kashima. All rights reserved.
//

#import "ViewController.h"
#import "R.h"

@interface ViewController ()

@end

@implementation ViewController
            
- (void)viewDidLoad {
    [super viewDidLoad];

    // Strings
    [self setTitle:[R stringTitleTop]];
    [self.message setText:[R stringLabelMessage]];

    // Colors
    [self.view setBackgroundColor:[R colorDefaultBg]];
    [self.message setTextColor:[R colorDefaultText]];

    // Drawables
    [self.image setImage:[R imageStar]];
}

@end
