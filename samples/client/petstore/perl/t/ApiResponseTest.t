# NOTE: This class is auto generated by the Swagger Codegen
# Please update the test case below to test the model.

use Test::More tests => 2;
use Test::Exception;

use lib 'lib';
use strict;
use warnings;


use_ok('WWW::OpenAPIClient::Object::ApiResponse');

my $instance = WWW::OpenAPIClient::Object::ApiResponse->new();

isa_ok($instance, 'WWW::OpenAPIClient::Object::ApiResponse');
