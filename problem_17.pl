#!/usr/bin/env perl
#
use warnings;
use strict;
use Lingua::EN::Numbers qw/num2en/;

my $sum = 0;

#for my $i (0 .. 1_000_000) {
for my $i (1 .. 1_000) {
  my $val = num2en($i);
  $val =~ s/[\s-]//g;
  $sum += length($val);
}

print "sum: $sum\n";
