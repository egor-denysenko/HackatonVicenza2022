import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { QuadriPage } from './quadri.page';

const routes: Routes = [
  {
    path: ':sightId',
    component: QuadriPage
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class QuadriPageRoutingModule {}
