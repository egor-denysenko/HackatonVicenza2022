import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { QuadriPage } from '../quadri/quadri.page';

import { StanzaPage } from './stanza.page';

const routes: Routes = [
  {
    path: ':roomId',
    component: StanzaPage
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class StanzaPageRoutingModule {}
