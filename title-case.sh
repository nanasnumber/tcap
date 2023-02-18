#!/bin/bash
src="source"

rm tcap.tmp

for i in `ls $src`;do
  grep '^## ' $src/$i/index.md >> tcap.tmp
done

while read i ;do

  title=${i//## /}
  tc=`tcap "$title"`
  echo ===========
  echo $title
  echo -----------
  echo $tc


  for i in `ls $src`;do
    sed -i "s/## $title/## $tc/g" $src/$i/index.md
  done

done < tcap.tmp
