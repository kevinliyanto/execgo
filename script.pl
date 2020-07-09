#!/usr/bin/perl -w

# I can probably do everything in one makefile but dealing with dirs are too annoying

$preset = shift @ARGV;

if (!$preset) {
  printf "no preset given. compile executor\n";
  system "cd executor && make && mv test ../test";
} elsif ($preset eq 'clean') {
  printf "cleaning all executables\n";
  my @exec = qw( test generate );
  for my $fi (@exec) {
    system "rm $fi";
  }
} elsif ($preset eq 'all') {
  printf "compile generator and executor\n";
  system "cd executor && make && mv test ../test";
  system "cd generator && make && mv generate ../generate";
} elsif ($preset eq 'generator') {
  printf "compile generator generator\n";
  system "cd generator && make && mv generate ../generate";
} else {
  die "unknown option $preset\nusage: $0 <clean|all|generator>\n"
}
