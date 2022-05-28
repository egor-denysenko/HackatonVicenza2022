import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-stanza',
  templateUrl: './stanza.page.html',
  styleUrls: ['./stanza.page.scss'],
})
export class StanzaPage implements OnInit {

  constructor(private router:Router) { }

  ngOnInit() {
  }

  VediamoStoQuadro(IdQuadro:number){
    this.router.navigate(['sight/'+ IdQuadro], {
      state: {
        IdQuadro: IdQuadro,
      },
    });
  }

}
