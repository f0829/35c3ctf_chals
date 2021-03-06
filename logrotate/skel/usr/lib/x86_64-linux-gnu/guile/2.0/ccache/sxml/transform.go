GOOF----LE-8-2.0l      ] ; 4   hÁ      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  sxml¤	g  	transform¤	 ¤		g  filenameS¤	
f  sxml/transform.scm¤	g  exportsS¤	g  SRV:send-reply¤	g  foldts¤	g  
post-order¤	g  pre-post-order¤	g  replace-range¤	 ¤	g  set-current-module¤	 ¤	 ¤	g  make-syntax-transformer¤	 ¤	 ¤	g  let*-values¤	g  macro¤	g  $sc-dispatch¤	 ¤	 ¤	g  _¤	g  any¤	¤	 g  syntax->datum¤	!  ¤	"  ¤	#g  datum->syntax¤	$# ¤	%# ¤	&g  begin¤	'g  let¤	(g  call-with-values¤	)g  lambda¤	*g  syntax-violation¤	+* ¤	,* ¤	-f  -source expression failed to match any pattern¤	.g  
procedure?¤	/g  display¤	0g  assq¤	1g  	*default*¤	2g  *text*¤	3g  
*preorder*¤	4g  *macro*¤	5g  append¤	6g  map¤	7g  error¤	8f  Unknown binding for ¤	9f   and no default¤	:g  reverse¤C 5  hÈ     ]4	
5 4 >  "  G   4"%&'()        hX   ó   ]
L L $  1 (      C    C    C ë       g  vars
		W g  initializer		W g  cont			W  g  filenamef  sxml/transform.scm
	V			X			Y		
	Y			X			W				[			Z			]			Z			^		 	^		%	^		0	`		G	\	 		W	   C  h(      - 1 3  (  C O  @            g  bindings
			# g  body			#  g  filenamef  sxml/transform.scm
	S
		T			T			T		!	b		#	U	 			#
   C       h   f   ]	4 5L 4?6^       g  args
		 g  v			  g  filenamef  sxml/transform.scm		S
 		   C,-   h(   d   ]	4 5$   O @ 6 \       g  y
		' g  tmp		'  g  filenamef  sxml/transform.scm
	S
 		'   C5R./      h°   ¥  ] (  C $   (    "ÿÿÜ &    "ÿÿÈ $   4L  5 "ÿÿ«4 5$  4 >   "  G    "ÿÿ4 >  "  G    "ÿÿc  "ÿÿV          g  	fragments
	 ¬ g  result	 ¬  g  filenamef  sxml/transform.scm
	n			o			q			o			r			o			r	%	&	r		*	s		.	o		1	s	&	:	s	 	=	t		>	t		B	o		E	u		F	u		K	u	$	O	u		W	u		X	v		]	v		_	v		c	o		d	w		g	w			l	w		w	x	 	x	 	z	 	z	 	z	 	{	 	{	 ¢	q	# ¬	q	 %	 ¬	  g  nameg  loop C        h    µ  -  1  3 O Q  6   ­      g  	fragments
			 g  loop		  g  filenamef  sxml/transform.scm
	d
		n	 			


  g  nameg  SRV:send-replyg  documentationf Output the @var{fragments} to the current output port.

The fragments are a list of strings, characters, numbers, thunks,
@code{#f}, @code{#t} -- and other fragments. The function traverses the
tree depth-first, writes out strings and characters, executes thunks,
and ignores @code{#f} and @code{'()}. The function returns @code{#t} if
anything was written at all; otherwise the result is @code{#f} If
@code{#t} occurs among the fragments, it is not written out but causes
the result of @code{SRV:send-reply} to be @code{#t}. CR012.034567892       hÐ     ] (  C $  ¢ $   4L 5$  "  L$  _$  F&   @&  4 ? "ÿÿ4 4L 55@4L 5@	6L 6L$  L
 6
	6           g  tree
	 Ë g  trigger	 « g  t		'	: g  binding		: «  g  filenamef  sxml/transform.scm
 		 			 		 		 		 		 		 		 ¡		 ¡	
	 ¢		' ¢		: ¡	
	B £		E ¦		F ¦		J £		L ¨		O ¨	 	T £		W ©		\ ©		^ ª		a ª		f £		g «		j «		o «		u «		x ­		| ®	  ®	"  ®	-  ®	5  ®	-  ®	  ­	  §	  §	-  §	7  §	- ¡ §	 ¥ ¥	 © ¥	6 « ¥	 ³ 	$ ¹ 	 ½ 	 Á 	 Å 	 Ç 	 É 	4 Ë 	 6	 Ë  g  nameg  loop C     hp     ]"4545$  "  $  45$  "  "  O Q  6           g  tree
		k g  bindings		k g  default-binding			k g  t			) g  text-binding		)	k g  text-handler		M	k g  loop		Z	k  g  filenamef  sxml/transform.scm
 
	 		 	 	 		 		 		 	!	 		 		) 		1 		2 		7 		9 		= 		@ 		G 	$	M 		Z 	 		k	  g  nameg  pre-post-order CRiR   h`   l  ](  C$  C"  +(  
64 5"ÿÿÕ4 5"ÿÿÁ6   d      g  fdown
		] g  fup		] g  fhere			] g  seed			] g  tree			] g  kid-seed			A g  kids			A  g  filenamef  sxml/transform.scm
 Å
	 Æ		 È			 Æ		 Ë		 Ì		& Í	
	' Î		4 Î	1	6 Î		9 Ï		A Î	
	A Ë		B Ë		M Ë	2	U Ë		] É	 		]	  g  nameg  foldts CR:5        hP    ]+ (  45D $  4L 5$   4455 "ÿÿ¹$  K4L$  "  >  G  $  
"   "ÿÿg  "ÿÿS4L5$  4 5 "ÿÿ-$  c4L$  "  >  G  $  "  $  $  
"  "   "ÿþÃ  "ÿþ²~      g  forest
	P g  keep?	P g  
new-forest		P g  node		P g  t		% ¯ g  node?		T  g  new-kids		p  g  keep?		p  g  t	 ¶P g  node?	 à? g  new-kids	 ü? g  keep?	 ü?  g  filenamef  sxml/transform.scm
 ü		 ý			 ý		 ý		 þ		 þ		 ÿ	
			% 		0		2		5	 	?		I		L		P 		S		$	T			T		W		_		b	#	j	7	o		z	 	 	( 	" 	 	 	 ¥	) ¯	 °	 ¶	 ¿	 Æ	. È	 Õ	 Ø	 Ü	 ß	$ à	 à	 ã	 ë	 î	# ö	7 û				%	  	# 	,& 	&/	?	B	P	 =	P	  g  nameg  loop C       h0     ]O  Q 4>  G C    ù       g  beg-pred
		, g  end-pred		, g  forest			, g  loop			, g  
new-forest		"	, g  keep?		"	,  g  filenamef  sxml/transform.scm
 ð
	%	$	%	4	!%	$	%%	 		,	  g  nameg  replace-range CRC  z       g  m
		(  g  filenamef  sxml/transform.scm		H
ø	d
¨ 
¯ ¶
 Å
Ä ð
 	Æ
   C6 