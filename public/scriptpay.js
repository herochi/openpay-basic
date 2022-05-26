var animation = bodymovin.loadAnimation({
  container: document.getElementById('animated-next'),
  renderer: 'svg',
  loop: false,
  autoplay: false,
  hideOnTransparent: true,
  path: 'https://api.myjson.com/bins/rwhs7.json'
})

$('#play').click( function() {
		animation.play();
});

$('#replay').click( function() {
		animation.goToAndStop(0);
		animation.play();
});

$('#next').click( function() {
  animation.playSegments([[34,115],[70,120],[70,160]],true);
$( "#animated-next" ).addClass( "move-icon" );
$( ".btn-text" ).addClass( "text-vanish" );
});