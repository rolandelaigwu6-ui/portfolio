const menuButton = document.querySelector('.menu-button');
const navigation = document.querySelector('.main-nav');

const brandStyles = document.createElement('link');
brandStyles.rel = 'stylesheet';
brandStyles.href = '/brand.css';
document.head.append(brandStyles);

const galleryStyles = document.createElement('link');
galleryStyles.rel = 'stylesheet';
galleryStyles.href = '/route-gallery.css';
document.head.append(galleryStyles);

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

const routeGalleries = {
  '/about/': {
    title: 'Leadership in focus',
    copy: 'A journey shaped by service, learning and an enduring connection to Benue South.',
    images: [['img2.jpeg', 'A leader rooted in community'], ['img3.jpeg', 'Service through relationships'], ['img4.jpeg', 'Experience with purpose']],
  },
  '/politics/': {
    title: 'Listening. Leading. Delivering.',
    copy: 'Leadership is strongest when it stays close to the people it serves.',
    images: [['img1.jpeg', 'Grassroots engagement'], ['img5.jpeg', 'Conversations that matter'], ['img6.jpeg', 'United by shared purpose']],
  },
  '/business/': {
    title: 'Relationships that create opportunity',
    copy: 'Enterprise grows through trust, collaboration and a long-term commitment to people.',
    images: [['img7.jpeg', 'Leadership that builds trust'], ['img8.jpeg', 'Creating meaningful connections'], ['img9.jpeg', 'Opportunity through partnership']],
  },
  '/philanthropy/': {
    title: 'Hope expressed through action',
    copy: 'Community support begins with presence, compassion and a commitment to restore dignity.',
    images: [['img10.jpeg', 'Compassion in action'], ['img11.jpeg', 'Standing with communities'], ['img12.jpeg', 'Hope for tomorrow']],
  },
  '/impact/': {
    title: 'Progress people can see',
    copy: 'Programmes gain meaning when communities are present, heard and strengthened.',
    images: [['img13.jpeg', 'Community impact'], ['img14.jpeg', 'Shared progress'], ['img15.jpeg', 'People first']],
  },
  '/news/': {
    title: 'From the field',
    copy: 'Official moments, public engagements and leadership stories in pictures.',
    images: [['img16.jpeg', 'Official engagement'], ['img17.jpeg', 'Leadership update'], ['img18.jpeg', 'In the community']],
  },
  '/events/': {
    title: 'Moments that bring people together',
    copy: 'Community gatherings, public conversations and events that strengthen shared purpose.',
    images: [['img19.jpeg', 'A community gathering'], ['img20.jpeg', 'Public occasion'], ['img21.jpeg', 'Shared moments']],
  },
  '/contacts/': {
    title: 'A leadership team that stays connected',
    copy: 'For invitations, media enquiries, partnership ideas and community feedback, the office is ready to listen.',
    images: [['img22.jpeg', 'Connect with the office'], ['img23.jpeg', 'Open dialogue'], ['img24.jpeg', 'Community first']],
  },
};

const gallery = routeGalleries[window.location.pathname];
if (gallery && !document.querySelector('.route-gallery')) {
  const figures = gallery.images.map(([image, caption], index) => `
    <figure>
      <img src="${assetBase}${image}" alt="Nelson Alapa — ${caption}" loading="${index === 0 ? 'eager' : 'lazy'}" />
      <figcaption>${caption}</figcaption>
    </figure>`).join('');
  document.querySelector('main')?.insertAdjacentHTML('beforeend', `
    <section class="route-gallery">
      <div class="container">
        <div class="route-gallery-heading"><h2>${gallery.title}</h2><p>${gallery.copy}</p></div>
        <div class="route-gallery-grid">${figures}</div>
      </div>
    </section>`);
}
