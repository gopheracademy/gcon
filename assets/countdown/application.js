var countdownTo = new Date(2017, 1-1, 1);

ready(function() {
  var countdownDaysEl = document.getElementById('js-countdown-days');
  var countdownHoursEl = document.getElementById('js-countdown-hours');
  var countdownMinutesEl = document.getElementById('js-countdown-minutes');
  var countdownSecondsEl = document.getElementById('js-countdown-seconds');

  var sec = 1000;
  var min = sec * 60;
  var hrs = min * 60;
  var day = hrs * 24;

  function update() {
    var nowMilis = Date.now();
    var targetMilis = countdownTo.getTime();
    var diffMilis = targetMilis - nowMilis;

    var daysLeft = Math.floor(diffMilis / day);
    var hoursLeft = Math.floor((diffMilis % day) / hrs);
    var minutesLeft = Math.floor((diffMilis % hrs) / min);
    var secondsLeft = Math.floor((diffMilis % min) / sec);

    if (diffMilis < 0) {
      daysLeft = 0;
      hoursLeft = 0;
      minutesLeft = 0;
      secondsLeft = 0;
    }

    countdownDaysEl.innerText = pad(daysLeft);
    countdownHoursEl.innerText = pad(hoursLeft);
    countdownMinutesEl.innerText = pad(minutesLeft);
    countdownSecondsEl.innerText = pad(secondsLeft);
  }
  update();
  setInterval(update, 1000);
});

function pad(num) {
  if (num > 99) {
    return num;
  }
  return ("00" + num).slice(-2);
}

function ready(fn) {
  if (document.readyState != 'loading') {
    fn();
  } else {
    document.addEventListener('DOMContentLoaded', fn);
  }
}
