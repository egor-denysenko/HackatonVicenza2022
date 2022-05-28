import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { QuadriPage } from '../quadri/quadri.page';
import { StanzaPage } from '../stanza/stanza.page';
import { HomePage } from './home.page';

const routes: Routes = [
  {
    path: '',
    component: HomePage,
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class HomePageRoutingModule {}
