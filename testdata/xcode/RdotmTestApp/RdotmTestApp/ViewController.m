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

    [self setTitle:[R string_title_top]];
    [self.message setText:[R string_label_message]];
}

@end
