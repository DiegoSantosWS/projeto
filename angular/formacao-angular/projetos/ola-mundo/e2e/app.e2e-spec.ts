import { AppPage } from './app.po';

describe('ola-mundo App', () => {
  let page: AppPage;

  beforeEach(() => {
    page = new AppPage();
  });

  it('should display welcome message', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('Olá Mundo!');
  });
});
