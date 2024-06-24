//top button
let topBtn = document.getElementById("topBtn");
window.addEventListener('scroll', () => {
    if (window.scrollY > 150) {
        topBtn.classList.toggle('hidden', false);
    } else {
        topBtn.classList.toggle('hidden', true);
    }
});
topBtn.addEventListener('click', () => {
    window.scrollTo({
        top: 0,
        behavior: 'smooth'
    });
});

//dark mode
const darkModeToggle = document.getElementById('darkModeToggle');
const sunIcon = document.getElementById('sunIcon');
const moonIcon = document.getElementById('moonIcon');
const localStorageTheme = localStorage.getItem('theme') || (window.matchMedia('(prefers-color-scheme: dark)').matches && 'dark');

const toggleDarkMode = (darkMode) => {
    document.documentElement.classList.toggle('dark', darkMode);
    localStorage.setItem('theme', darkMode ? 'dark' : 'light');
    sunIcon.classList.toggle('hidden', !darkMode);
    moonIcon.classList.toggle('hidden', darkMode);
};

darkModeToggle.checked = localStorageTheme === 'dark';
toggleDarkMode(darkModeToggle.checked);

darkModeToggle.addEventListener('change', () => toggleDarkMode(darkModeToggle.checked));

// custom cursor
let innerCursor = document.querySelector(".inner-cursor");
let outerCursor = document.querySelector(".outer-cursor");
document.addEventListener("mousemove", moveCursor);
function moveCursor(e) {
    let x = e.clientX;
    let y = e.clientY;
    innerCursor.style.left = `${x}px`;
    innerCursor.style.top = `${y}px`;
    outerCursor.style.left = `${x}px`;
    outerCursor.style.top = `${y}px`;
}
let elements = Array.from(document.querySelectorAll("a, p, span, h1, h2, h3, h4, h5, h6, li, td, th, button, input, textarea"));
elements.forEach((element) => {
    element.addEventListener("mouseover", () => {
        innerCursor.classList.add("grow");
        outerCursor.classList.add("grow");
    });
    element.addEventListener("mouseleave", () => {
        innerCursor.classList.remove("grow");
        outerCursor.classList.remove("grow");
    });
});

// post sidebar
document.addEventListener("DOMContentLoaded", function () {
    const content = document.getElementById("content");
    const sidebar = document.getElementById("post-sidebar");
    const headings = content.querySelectorAll("h1, h2, h3, h4, h5, h6");

    const headingStructure = [];
    const levelStack = [];

    headings.forEach(heading => {
        const level = parseInt(heading.tagName[1]);
        const text = heading.innerText;
        const id = text.toLowerCase().replace(/\s+/g, '-');
        heading.setAttribute('id', id);

        const newItem = { id, text, level, children: [] };

        while (levelStack.length && levelStack[levelStack.length - 1].level >= level) {
            levelStack.pop();
        }

        if (levelStack.length === 0) {
            headingStructure.push(newItem);
        } else {
            levelStack[levelStack.length - 1].children.push(newItem);
        }

        levelStack.push(newItem);
    });

    function createList(items) {
        const ul = document.createElement('ul');
        ul.classList.add('list-none', 'pl-4');
        items.forEach(item => {
            const li = document.createElement('li');
            li.innerHTML = `<a href="#${item.id}" class="hover:underline">${item.text}</a>`;
            const childUl = createList(item.children);
            childUl.classList.add('hidden');
            li.appendChild(childUl);
            if (item.children.length > 0) {
                li.classList.add('has-children', 'cursor-pointer');
                li.addEventListener('click', function (event) {
                    event.stopPropagation();
                    this.classList.toggle('open');
                    childUl.classList.toggle('hidden');
                });
            }
            ul.appendChild(li);
        });
        return ul;
    }

    sidebar.appendChild(createList(headingStructure));

    const links = sidebar.querySelectorAll('a');

    links.forEach(link => {
        link.addEventListener('click', function (event) {
            event.preventDefault();
            const section = document.querySelector(this.getAttribute('href'));
            section.scrollIntoView({ behavior: 'smooth' });
        });
    });

    let activeLink = null;
    function setActiveLink() {
        links.forEach(link => {
            link.classList.remove('text-lime-600', 'dark:text-lime-300');
        });
        for (let i = links.length - 1; i >= 0; i--) {
            const section = document.querySelector(links[i].getAttribute('href'));
            const sectionTop = section.getBoundingClientRect().top;
            if (sectionTop <= window.innerHeight / 2) {
                activeLink = links[i];
                break;
            }
        }

        if (activeLink) {
            activeLink.classList.add('text-lime-600', 'dark:text-lime-300');
            const parentLi = activeLink.closest('li');
            const childUl = parentLi.querySelector('ul');
            if (childUl) {
                childUl.classList.remove('hidden');
            }

            const siblings = Array.from(parentLi.parentNode.children).filter(el => el !== parentLi);
            siblings.forEach(sibling => {
                const siblingChildUl = sibling.querySelector('ul');
                if (siblingChildUl) {
                    siblingChildUl.classList.add('hidden');
                }
            });
        }
    }

    document.addEventListener('scroll', setActiveLink, { passive: true });
    setActiveLink();
});

// set active nav-link on click and remove active class from other nav-links
const navLinks = document.querySelectorAll('.nav-link');
navLinks.forEach((navLink, index) => {
    navLink.addEventListener('click', (event) => {
        if (navLink.getAttribute('target') === '_blank') return;

        navLinks.forEach(link => link.classList.remove('active'));
        navLink.classList.add('active');
        localStorage.setItem('activeNavLink', index); // store active link index
    });
});

window.onload = () => {
    let activeNavLinkIndex = localStorage.getItem('activeNavLink');
    if (activeNavLinkIndex === null) {
        activeNavLinkIndex = 0; // default to first link (Home)
    }
    navLinks[activeNavLinkIndex].classList.add('active');
};