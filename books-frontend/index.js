import { LitElement, css, html } from 'lit';

export class SimpleGreeting extends LitElement {
  static get properties() {
    return {
      books: [],
      loading: false,
      error: false,
      showingList: true,
      cart: [],
      cartSize: 0,
    };
  }

  static get styles() {
    return css`
      :host {
        height: 100%;
        display: flex;
        flex-direction: column;
      }

      nav {
        display: flex; /* 1 */
        justify-content: space-between; /* 2 */
        padding: 1rem 2rem; /* 3 */
        background: #cfd8dc; /* 4 */
      }
      
      nav ul {
        display: flex; /* 5 */
        list-style: none; /* 6 */
      }
      
      nav li {
        padding-left: 1rem; /* 7! */
      }
      nav a {
        text-decoration: none;
        color: #0d47a1
      }
      /* 
        Extra small devices (phones, 600px and down) 
      */
      @media only screen and (max-width: 600px) {
        nav {
          flex-direction: column;
        }
        nav ul {
          flex-direction: column;
          padding-top: 0.5rem;
        }
        nav li {
          padding: 0.5rem 0;
        }
      }
      article {
        --img-scale: 1.001;
        --title-color: black;
        --link-icon-translate: -20px;
        --link-icon-opacity: 0;
        position: relative;
        border-radius: 16px;
        box-shadow: none;
        background: #fff;
        transform-origin: center;
        transition: all 0.4s ease-in-out;
        overflow: hidden;
      }
      
      article a::after {
        position: absolute;
        inset-block: 0;
        inset-inline: 0;
        cursor: pointer;
        content: "";
      }
      
      /* basic article elements styling */
      article h2 {
        margin: 0 0 18px 0;
        font-family: "Bebas Neue", cursive;
        font-size: 1.9rem;
        letter-spacing: 0.06em;
        color: var(--title-color);
        transition: color 0.3s ease-out;
      }
      
      figure {
        margin: 0;
        padding: 0;
        aspect-ratio: 16 / 9;
        overflow: hidden;
      }
      
      article img {
        max-width: 100%;
        transform-origin: center;
        transform: scale(var(--img-scale));
        transition: transform 0.4s ease-in-out;
      }
      
      .article-body {
        padding: 24px;
      }
      
      article a {
        display: inline-flex;
        align-items: center;
        text-decoration: none;
        color: #28666e;
      }
      
      article a:focus {
        outline: 1px dotted #28666e;
      }
      
      article a .icon {
        min-width: 24px;
        width: 24px;
        height: 24px;
        margin-left: 5px;
        transform: translateX(var(--link-icon-translate));
        opacity: var(--link-icon-opacity);
        transition: all 0.3s;
      }
      
      /* using the has() relational pseudo selector to update our custom properties */
      article:has(:hover, :focus) {
        --img-scale: 1.1;
        --title-color: #28666e;
        --link-icon-translate: 0;
        --link-icon-opacity: 1;
        box-shadow: rgba(0, 0, 0, 0.16) 0px 10px 36px 0px, rgba(0, 0, 0, 0.06) 0px 0px 0px 1px;
      }
      
      
      /************************ 
      Generic layout (demo looks)
      **************************/      
      .articles {
        display: grid;
        max-width: 1200px;
        margin-inline: auto;
        padding-inline: 24px;
        grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
        gap: 24px;
        padding-top: 12px;
        padding-bottom: 12px;
        background-color: rgb(38, 48, 56);
      }
      
      @media screen and (max-width: 960px) {
        article {
          container: card/inline-size;
        }
        .article-body p {
          display: none;
        }
      }
      
      @container card (min-width: 380px) {
        .article-wrapper {
          display: grid;
          grid-template-columns: 100px 1fr;
          gap: 16px;
        }
        .article-body {
          padding-left: 0;
        }
        figure {
          width: 100%;
          height: 100%;
          overflow: hidden;
        }
        figure img {
          height: 100%;
          aspect-ratio: 1;
          object-fit: cover;
        }
      }
      
      .sr-only:not(:focus):not(:active) {
        clip: rect(0 0 0 0); 
        clip-path: inset(50%);
        height: 1px;
        overflow: hidden;
        position: absolute;
        white-space: nowrap; 
        width: 1px;
      }
      .loader {
        width: 48px;
        height: 48px;
        border: 5px solid #FFF;
        border-bottom-color: transparent;
        border-radius: 50%;
        display: inline-block;
        box-sizing: border-box;
        animation: rotation 1s linear infinite;
      }
    
      @keyframes rotation {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
      }
      .wrapper {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        flex: 1;
      }

      .error {
        color: white;
      }

      h3 {
        color: white;
      }

      .checked {
        background-color: blanchedalmond;
      }
      .float{
        display: flex;
        flex-direction: column;
        justify-content: center;
        font-weight: bold;
        position:fixed;
        width:60px;
        height:60px;
        bottom:40px;
        right:40px;
        background-color:#0C9;
        color:#FFF;
        border-radius:50px;
        text-align:center;
        box-shadow: 2px 2px 3px #999;
      }
      
      .my-float{
        margin-top:22px;
      }
    `;
  }

  constructor() {
    super();

    this.cart = [];
    this.books = [];
    this.loading = true;
    this.error = false;
    this.showingList = true;
    this.cartSize = 0;
  }

  connectedCallback() {
    super.connectedCallback();

    this.getBooksList();
  }

  getBooksList() {
    this.loading = true;
    this.error = false;

    fetch('/api/books')
      .then(response => response.json())
      .then(books => {
        this.books = books;
        this.loading = false;
      }).catch(() => {
        this.error = true;
        this.loading = false;
      });
    
    this.cart = [];
    this.cartSize = 0;
  }

  showList(event) {
    event.preventDefault();

    this.showingList = true;
    this.getBooksList();
  }

  getBooksPurchased() {
    this.loading = true;
    this.error = false;

    fetch('/api/bookstore/purchased')
      .then(response => response.json())
      .then(books => {
        this.books = books;
        this.loading = false;
      }).catch(() => {
        this.error = true;
        this.loading = false;
      });

    this.cart = [];
    this.cartSize = 0;
  }

  showPurchased(event) {
    event.preventDefault();

    this.showingList = false;
    this.getBooksPurchased();
  }

  addToCart(event) {
    event.preventDefault();

    const foundIndex = this.cart.findIndex(title => title === event.currentTarget.dataset.title);

    if (foundIndex < 0) {
      this.cart.push(event.currentTarget.dataset.title);
      this.cartSize++;
    } else {
      this.cart.splice(foundIndex, 1);
      this.cartSize--;
    }
    event.currentTarget.classList.toggle('checked')
  }

  buyBooks(event) {
    event.preventDefault();
    if (this.cart.length > 0) {
      this.loading = true;
      this.error = false;
      fetch('/api/bookstore/purchase', {
        method: 'POST',
        body: JSON.stringify(this.cart),
      })
        .then(() => this.loading = false)
        .catch(() => {
          this.loading = false;
          this.error = true;
        });
    }

    this.cart = [];
    this.cartSize = 0;
  }

  render() {
    return html`
      <nav>
        <h2>Books</h2>
        <ul>
          <li><a @click="${this.showList}" href="#">List</a></li>
          <li><a @click="${this.showPurchased}" href="#">Purchased</a></li>
        </ul>
      </nav>
      ${this.loading ?
        html`
          <div class="wrapper">
            <span class="loader"></span>
          </div>
        `
        : html`
        <section class="articles">
          ${this.showingList ? html`<h3>List</h3>` : html`<h3>Purchased</h3>`}
          ${!this.error ?
            this.books.map(book => {
            return html`
              <article data-title="${book.title}" @click="${this.addToCart}">
                <div class="article-wrapper">
                  <figure>
                    <img src="${book.cover}" alt="Book cover" />
                  </figure>
                  <div class="article-body">
                    <h2>${book.title}</h2>
                    <p>
                      ${book.author}
                    </p>
                    <p>
                      ${book.description}
                    </p>
                  </div>
                </div>
              </article>
            `
          }) : html`<p class="error">Something happened!</p>`}
        </section>
        `
      }
      <a href="#" class="float" @click="${this.buyBooks}">
        ${this.cart.length}
      </a>
      `

  }
}

customElements.define('simple-greeting', SimpleGreeting);