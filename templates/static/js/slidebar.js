function slideBar() {
    var menuBtn = document.querySelector('.menu-btn');
    var nav = document.querySelector('nav');
    var lineOne = document.querySelector('nav .menu-btn .line--1');
    var lineTwo = document.querySelector('nav .menu-btn .line--2');
    var lineThree = document.querySelector('nav .menu-btn .line--3');
    var link = document.querySelector('nav .nav-links');
    menuBtn.addEventListener('click', () => {
        nav.classList.toggle('nav-open');
        lineOne.classList.toggle('line-cross');
        lineTwo.classList.toggle('line-fade-out');
        lineThree.classList.toggle('line-cross');
        link.classList.toggle('fade-in');
    })
    
    $(window).scroll(function() {
        var scroll = $(window).scrollTop();
        //console.log(scroll);
        if (scroll >= 1) {
            nav.classList.remove('nav-open');
            lineOne.classList.remove('line-cross');
            lineTwo.classList.remove('line-fade-out');
            lineThree.classList.remove('line-cross');
            link.classList.remove('fade-in');
        } else {
    
        }
    });
}

slideBar()