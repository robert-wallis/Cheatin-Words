Abandoned
=========

This project is abandoned because app engine started charging me and there was no traffic to the site.  However the repo is still here for your amusement.

Cheatin' Words
==============

This is my first Go Lanugage (golang) project on the Google App Engine.  I made it to have a real project that I could test to see how much I like go as a web development language.

Previously at <http://words.ueuno.com/>

After burning through a few permutations implementations and finally transcribing python's itertools implementation, I finally got the site runnable.

It is dreadfully slow.  I am probably not writing the go code in the most optimal way.  If you can fix it, please request a pull on github <https://github.com/robert-wallis/Cheatin-Words>.  My confidence in go has been shattered by these performance issues coupled with the pain it took manually writing out things I took for granted in Python.  I can't motivate myself to implement wildcards, scoring, and forced letters in go because it seems too daunting.  I've quickly sketched it out here:

`def findWildcard(self, letters, requiredLetters, minLength=2):
    return sorted([w for l in range(ord('a'), ord('z')+1) for w in s.findPermutations(letters+chr(l)) if w in s.words and requiredLetters in w and len(w) >= minLength ], key=self.score)`

Anyway, have fun with the code, if you can make it run at about 100ms on Google App Engine, I'll come back to go.  I love many things about go.

