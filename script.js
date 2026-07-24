const menuButton = document.querySelector('.menu-button');
const navigation = document.querySelector('.main-nav');

const brandStyles = document.createElement('link');
brandStyles.rel = 'stylesheet';
brandStyles.href = '/brand.css';
document.head.append(brandStyles);

const favicon = document.createElement('link');
favicon.rel = 'icon';
favicon.type = 'image/jpeg';
favicon.href = '/assets/images/logo.jpeg';
document.head.append(favicon);

document.querySelectorAll('.brand .seal').forEach((seal) => {
  seal.innerHTML = '<img src="/assets/images/logo.jpeg" alt="" />';
});

const foundationLink = navigation?.querySelector('a[href="/philanthropy/"]');
if (foundationLink) {
  foundationLink.textContent = 'Foundation';
  if (!navigation.querySelector('a[href="/impact/"]')) {
    const impactLink = document.createElement('a');
    impactLink.href = '/impact/';
    impactLink.textContent = 'Impact';
    foundationLink.insertAdjacentElement('afterend', impactLink);
  }
}

menuButton?.addEventListener('click', () => {
  const open = navigation.classList.toggle('open');
  menuButton.setAttribute('aria-expanded', String(open));
  menuButton.textContent = open ? '×' : '☰';
});

navigation?.querySelectorAll('a').forEach((link) => link.addEventListener('click', () => {
  navigation.classList.remove('open');
  menuButton?.setAttribute('aria-expanded', 'false');
  if (menuButton) menuButton.textContent = '☰';
}));

const assetBase = '/assets/images/';
const pageImages = {
  '/': 'img2.jpeg',
  '/about/': 'img2.jpeg',
  '/politics/': 'img5.jpeg',
  '/philanthropy/': 'img10.jpeg',
  '/impact/': 'img13.jpeg',
  '/news/': 'img16.jpeg',
  '/events/': 'img19.jpeg',
  '/contacts/': 'img22.jpeg',
};

const imageURL = pageImages[window.location.pathname];
if (imageURL) {
  const hero = document.querySelector('.hero');
  const routePhoto = document.querySelector('.route-photo');
  if (hero) hero.style.setProperty('background-image', `linear-gradient(105deg,#15384b 8%,rgba(21,56,75,.69) 43%,rgba(21,56,75,.2)),url('${assetBase}${imageURL}')`, 'important');
  if (routePhoto) routePhoto.style.setProperty('background-image', `linear-gradient(90deg,rgba(16,44,62,.25),rgba(16,44,62,.25)),url('${assetBase}${imageURL}')`, 'important');
}

const cardImages = {
  '.image-road': 'img1.jpeg',
  '.image-school': 'img2.jpeg',
  '.image-clinic': 'img3.jpeg',
};

Object.entries(cardImages).forEach(([selector, image]) => {
  const element = document.querySelector(selector);
  if (element) element.style.backgroundImage = `url('${assetBase}${image}')`;
});
