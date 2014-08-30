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
    [self setTitle:[R string_title_top]];
    [self.message setText:[NSString stringWithFormat:[R string_label_message], @"world"]];
    [self.message2 setText:[NSString stringWithFormat:[R string_label_message2],
                            @"foo",
                            [R integer_sample_number]]];

    // Colors
    [self.view setBackgroundColor:[R color_default_bg]];
    [self.message setTextColor:[R color_default_text]];

    // Drawables
    [self.image setImage:[R drawable_star]];

    // Integer array
    NSMutableString *message3 = [NSMutableString stringWithString:@""];
    for (NSNumber *i in [R array_integer_foobar]) {
        [message3 appendFormat:@"%@ ", i];
    }
    [self.message3 setText:message3];
}

@end
