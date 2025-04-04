// Code generated by astool. DO NOT EDIT.

package streams

import (
	typealbum "codeberg.org/superseriousbusiness/activity/streams/impl/funkwhale/type_album"
	typeartist "codeberg.org/superseriousbusiness/activity/streams/impl/funkwhale/type_artist"
	typelibrary "codeberg.org/superseriousbusiness/activity/streams/impl/funkwhale/type_library"
	typetrack "codeberg.org/superseriousbusiness/activity/streams/impl/funkwhale/type_track"
	vocab "codeberg.org/superseriousbusiness/activity/streams/vocab"
)

// FunkwhaleAlbumIsDisjointWith returns true if Album is disjoint with the other's
// type.
func FunkwhaleAlbumIsDisjointWith(other vocab.Type) bool {
	return typealbum.AlbumIsDisjointWith(other)
}

// FunkwhaleArtistIsDisjointWith returns true if Artist is disjoint with the
// other's type.
func FunkwhaleArtistIsDisjointWith(other vocab.Type) bool {
	return typeartist.ArtistIsDisjointWith(other)
}

// FunkwhaleLibraryIsDisjointWith returns true if Library is disjoint with the
// other's type.
func FunkwhaleLibraryIsDisjointWith(other vocab.Type) bool {
	return typelibrary.LibraryIsDisjointWith(other)
}

// FunkwhaleTrackIsDisjointWith returns true if Track is disjoint with the other's
// type.
func FunkwhaleTrackIsDisjointWith(other vocab.Type) bool {
	return typetrack.TrackIsDisjointWith(other)
}
